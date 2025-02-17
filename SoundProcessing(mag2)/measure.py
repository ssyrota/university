import difflib
import nltk


class SpeachRecognitionMeasure:
    def __init__(self, ref, hyp) -> None:
        self.ref = self._preprocess(ref)
        self.hyp = self._preprocess(hyp)

    def wer(self):
        compared_words = self._compare_words()
        return (len(compared_words['replacements']) + len(compared_words['insertions']) + len(compared_words['deletions']))/len(self.ref.split())

    def _compare_words(self):
        return self._diff(self.ref.split(), self.hyp.split())

    def cer(self):
        compared_chars = self._compare_chars()
        return (len(compared_chars['replacements']) + len(compared_chars['insertions']) + len(compared_chars['deletions']))/len(self.ref)

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
            if tag == 'replace':
                ref_part = ref[i1:i2]
                hyp_part = hyp[j1:j2]
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
                print("deletions: ", ref[i1:i2])
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
                     "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ь", "ю", "я", "-", "’"]

    def _preprocess(self, text: str) -> str:
        cleaned_text = text.lower().replace("'", "’").strip().replace(
            "?", ".").replace("!", ".").replace(",", "")

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

    def to_file(self, filename: str):
        with open(filename, 'w') as f:
            f.write("# WER\n\n")
            f.write(self.markdown_wer())
            f.write("\n\n")
            f.write("# CER\n\n")
            f.write(self.markdown_cer())


ref = 'Привіт світ, а що означає привіт для великих мовних моделей?'
hyp = 'Привіт світ, що означає пивіт для дуже великих мовних моделей, знаєш?'

measure = SpeachRecognitionMeasure(ref, hyp)
measure.to_file('report.md')
