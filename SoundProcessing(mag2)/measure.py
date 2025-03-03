import difflib
import nltk
import pandas as pd


class SpeachRecognitionMeasure:
    def __init__(self, ref, hyp) -> None:
        self.ref = self._preprocess(ref)
        self.hyp = self._preprocess(hyp)

    def wer(self):
        compared_words = self._compare_words()
        replacements = self._flatten(
            [i[0] for i in compared_words['replacements']])
        insertions = self._flatten([i for i in compared_words['insertions']])
        deletions = self._flatten([i for i in compared_words['deletions']])
        return (len(replacements) + len(insertions) + len(deletions))/len(self.ref.split())

    def cer(self):
        compared_chars = self._compare_chars()
        replacements = self._flatten(
            [i[0] for i in compared_chars['replacements']])
        insertions = self._flatten([i for i in compared_chars['insertions']])
        deletions = self._flatten([i for i in compared_chars['deletions']])
        return (len(replacements) + len(insertions) + len(deletions))/len(self.ref)

    def _compare_words(self):
        return self._diff(self.ref.split(), self.hyp.split())

    def _compare_chars(self):
        return self._diff(self.ref, self.hyp)

    def _diff(self, ref, hyp):
        diff = difflib.SequenceMatcher(None, ref, hyp)
        replacements = []
        insertions = []
        deletions = []

        ref_compared = []
        hyp_compared = []
        TRUELY_RECOGNIZED_TOKEN = "*"
        INSERTION_TOKEN = "+"
        DELETION_TOKEN = "-"
        for tag, i1, i2, j1, j2 in diff.get_opcodes():
            print("tag: ", tag)
            print("i1: ", i1)
            print("i2: ", i2)
            print("j1: ", j1)
            print("j2: ", j2)
            if tag == 'replace':
                ref_part = ref[i1:i2]
                hyp_part = hyp[j1:j2]

                if len(hyp_part) < len(ref_part):
                    missing = len(ref_part) - len(hyp_part)
                    if isinstance(hyp_part, list):
                        hyp_part = hyp_part + \
                            [DELETION_TOKEN for _ in range(missing)]
                    else:
                        hyp_part = hyp_part + \
                            DELETION_TOKEN * missing
                if len(ref_part) < len(hyp_part):
                    missing = len(hyp_part) - len(ref_part)
                    if isinstance(ref_part, list):
                        ref_part = ref_part + \
                            [DELETION_TOKEN for _ in range(missing)]
                    else:
                        ref_part = ref_part + \
                            DELETION_TOKEN * missing

                replacements.append((ref_part, hyp_part))
                ref_compared.append(ref_part)
                hyp_compared.append(hyp_part)
            elif tag == 'insert':
                insertions.append(hyp[j1:j2])
                ref_compared.append(INSERTION_TOKEN*len(hyp[j1:j2]))
                hyp_compared.append(hyp[j1:j2])
            elif tag == 'delete':
                deletions.append(ref[i1:i2])
                ref_compared.append(ref[i1:i2])
                hyp_compared.append(DELETION_TOKEN*len(ref[i1:i2]))
            elif tag == 'equal':
                ref_compared.append(ref[i1:i2])
                hyp_compared.append(TRUELY_RECOGNIZED_TOKEN * len(ref[i1:i2]))

        ref_compared = self._flatten(ref_compared)
        hyp_compared = self._flatten(hyp_compared)

        if len(ref_compared) != len(hyp_compared):
            print("ref_compared: ", ref_compared)
            print("hyp_compared: ", hyp_compared)
            raise ValueError(
                f"ref_compared and hyp_compared have different lengths: ref_compared: {len(ref_compared)} != hyp_compared: {len(hyp_compared)}")

        return {'replacements': replacements, 'insertions': insertions, 'deletions': deletions, 'ref_compared': ref_compared, 'hyp_compared': hyp_compared}

    def _flatten(self, list_of_lists):
        return [item for sublist in list_of_lists for item in sublist]

    tokenizer = nltk.SpaceTokenizer()
    allowed_chars = ["а", "б", "в", "г", "ґ", "д", "е", "є", "ж", "з", "и", "і", "ї", "й", "к", "л",
                     "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ь", "ю", "я", "’"]

    def _preprocess(self, text: str) -> str:
        cleaned_text = text.lower().replace("'", "’").strip().replace(
            "?", ".").replace("!", ".").replace(",", "").replace(".", "")

        cleaned_text = cleaned_text.split(".")
        out_text = []
        for text in cleaned_text:
            text = text.strip()
            words = [i for i in self.tokenizer.tokenize(
                text) if not i.isdigit()]
            words = [self.clear_word(w) for w in words]
            words = [w for w in words if w != ""]
            if all([len(i) <= 1 for i in words]):
                continue
            if len(words) == 0:
                continue
            out_text.append(
                " ".join(words))
        cleaned_text = "\n".join(out_text)
        return cleaned_text

    def clear_word(self, word: str) -> str:
        if all([i in self.allowed_chars for i in word]):
            return word
        return ""

    def csv_to_file(self, wer_filename: str, cer_filename: str):
        print("WER")
        with open(wer_filename, 'w') as f:
            f.write(self.csv_wer())
        print("CER")
        with open(cer_filename, 'w') as f:
            f.write(self.csv_cer())

    def csv_cer(self):
        cer = self.cer()
        compared_chars = self._compare_chars()
        df = pd.DataFrame({
            'ref': compared_chars['ref_compared'],
            'hyp': compared_chars['hyp_compared'],
            'cer': cer
        })
        csv_content = df.to_csv(index=False).strip()
        return csv_content

    def csv_wer(self):
        wer = self.wer()
        compared_words = self._compare_words()
        df = pd.DataFrame({
            'ref': compared_words['ref_compared'],
            'hyp': compared_words['hyp_compared'],
            'wer': wer
        })
        return df.to_csv(index=False).strip()

    def markdown_wer(self):
        wer = self.wer()
        compared_words = self._compare_words()
        ref = compared_words['ref_compared']
        hyp = compared_words['hyp_compared']
        return self._markdown_report("WER", wer, ref, hyp)

    def markdown_cer(self):
        cer = self.cer()
        compared_chars = self._compare_chars()
        ref = compared_chars['ref_compared']
        hyp = compared_chars['hyp_compared']
        return self._markdown_report("CER", cer, ref, hyp)

    def _markdown_report(self, metric_name: str, metric_value: float, ref, hyp):
        header = f"| ref | hyp | {metric_name}: {metric_value} |"
        separator = "| --- | --- | --- |"
        rows = [f"| {ref} | {hyp} | |" for ref, hyp in zip(ref, hyp)]
        return "\n".join([header, separator, *rows])

    def md_to_file(self, filename: str):
        with open(filename, 'w') as f:
            report_str = f"""# WER

{self.markdown_wer()}

# CER

{self.markdown_cer()}
"""
            f.write(report_str)


ref = """Сидить батько кінець стола, На руки схилився, Не дивиться на світ Божий: Тяжко зажурився. Коло його стара мати Сидить на ослоні, За сльозами ледве-ледве Вимовляє доні: «Що весілля, доню моя? А де ж твоя пара? Де світилки з друженьками, Старости, бояре? В Московщині, доню моя! Іди ж їх шукати, Та не кажи добрим людям, Що є в тебе мати. Проклятий час-годинонька, Що ти народилась! Якби знала, до схід сонця Була б утопила... Здалась тоді б ти гадині, Тепер — москалеві... Доню моя, доню моя, Цвіте мій рожевий! Як ягодку, як пташечку, Кохала, ростила На лишенько... Доню моя, Що ти наробила?.. Оддячила!.. Іди ж, шукай У Москві свекрухи. Не слухала моїх річей, То її послухай"""
hyp_wav2_vec = """сидить батько кінець стола на руки схилився не дивиться на світбоже тяжко зажурився коло його стара мати сидить на ослоні за сльозами ледве ледве вимовляє доні що весілля доню моє а де ж твоя пара де світлинки з друженьками старости бо яри в московщині доню моє іди ж їх шукати та не кажи добрим людям що є в тебе мат проклятий час годинонька що ти народилась якби знала до схід сонця була б утопила здалась тобі б ти гадині тепер москалеві доню моє доню моє цвітам ій рожевий як ягодку як пташичку кохала ростила на лишенько доню моє що ти наробила отдячила ідиж шукай у москві свекрухи не слухала моїх річей то її послухай"""
hyp_deep_speech = """сидить батько кінець столу на руки схилився не дивиться на світ божий тяжко зажурився коло його стара мати сидить на ослоні оловоносних моноклональних"""

hyp_espnet_wav2_vec = """сидить батько кінець стола на руки схвилився не дивиться на світ божей тяжко зажурився коло його стара мати сидить на ослоні за сльозами ледвали две вимовляє доні що весілля доню моя а де ж твоя пара де світилки з друженьками старости бо яри восковщині доню моя іди ж їх шукати та не кажи добрим людям що є в тебе мати проклятий час годинонька що ти народилась якби знала до схід сонця була б утопила здалась тоді пти гадені тепер м москалеві доню моя доню моя світемій рожевий як я готку як пташичку кохала ростила на лишенько доню моє що ти не робила отдячила іди шукав як ву москві свокрухи не слухала моїх річей то їй послухай"""
hyp_styletts2_wav2_vec = """сидить батько кінець столе на руки схилився не дивиться на світ божий тяжко зажурився коло його стара мати сидить на ослоні за сльозами ледве ледве вимовляє доні що весілля доню моє а де ж твоя паре де світилки з друженьками старости бояри в московщині доню моє іди ж їх шукати та не кажи добрим людям що є в тебе мати проклятий час з годин оньке що ти народилась як би знала до схід сонця була бутопила здалась тоді пти гади ні тепер москалеві доню моя доню моя цвітемій рожевий як яготку як пташечку кохала розстила на лишенько доню моє що ти наробила отдячила іди ж шукаю москві свекрухи не слухала моїх річей то її послухай"""

measure = SpeachRecognitionMeasure(ref, hyp_styletts2_wav2_vec)
measure.md_to_file('report.md')
measure.csv_to_file('wer.csv', 'cer.csv')

measure = SpeachRecognitionMeasure(ref, hyp_espnet_wav2_vec)
measure.md_to_file('espnet_report.md')
measure.csv_to_file('espnet_wer.csv', 'espnet_cer.csv')
