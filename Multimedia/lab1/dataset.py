import pandas as pd
import librosa

class UrbanSound:
  def __init__(self, meta_path:str, audio_path:str):
    self.audio_path = audio_path
    self.data = pd.read_csv(meta_path)

  def generate_images(self):
    librosa