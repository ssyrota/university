import librosa
import librosa.display
import numpy as np
import matplotlib.pyplot as plt
# memory leak fix
import matplotlib
matplotlib.use('Agg')


class Sound:
    def __init__(self, path):
        self.sound_path = path


    def export_melgram(self, export_path):
        MELS=128
        samples, sample_rate = librosa.load(self.sound_path, sr=None)
        sgram = librosa.stft(samples)
        sgram_mag, _ = librosa.magphase(sgram)
        mel_scale_sgram = librosa.feature.melspectrogram(S=sgram_mag, sr=sample_rate, n_mels=MELS)
        mel_sgram = librosa.amplitude_to_db(mel_scale_sgram, ref=np.min)
        
        # hack to make each mel have one pixel of height
        fig, ax = plt.subplots(figsize=(256 / 100, 2*MELS / 100))
        librosa.display.specshow(mel_sgram, ax=ax)
        ax.axis('off')
        plt.tight_layout(pad=0)
        plt.savefig(export_path, format='jpg', bbox_inches='tight', pad_inches=0)
        # memory leak fix
        plt.close(fig)
        fig.clf()
        del fig, ax
