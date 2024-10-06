import librosa
import librosa.display
import numpy as np
import matplotlib.pyplot as plt

class Sound:
    def __init__(self, path):
        self.sound_path = path

    def export_merlgram(self, export_path):
        samples, sample_rate = librosa.load(self.sound_path, sr=None)
        sgram = librosa.stft(samples)
        sgram_mag, _ = librosa.magphase(sgram)
        mel_scale_sgram = librosa.feature.melspectrogram(S=sgram_mag, sr=sample_rate)
        mel_sgram = librosa.amplitude_to_db(mel_scale_sgram, ref=np.min)
        librosa.display.specshow(mel_sgram)
        plt.savefig(export_path)