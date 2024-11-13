from os import mkdir, path
import pandas as pd
from sound import Sound


class UrbanSound8K:
    def __init__(self, dataset_path: str, meta_path: str, audio_path: str):
        self.audio_path = audio_path
        self.images_path = path.join(dataset_path, "images")
        self.data = pd.read_csv(meta_path)

    def folds(self):
        folds = []
        for i in range(1, 11):
            fold = path.join(self.images_path, "fold"+str(i))
            folds.append(fold)
        return folds


    def generate_images(self):
        self._ensure_image_folders()
        for i in range(len(self.data)):
            row = self.data.iloc[i]
            sound_file_path = path.join(
                self.audio_path, "fold"+str(row['fold']), row['slice_file_name'])
            jpeg_folder = path.join(self.images_path, "fold"+str(row['fold']), row['class'])
            self._ensure_folder(jpeg_folder)
            jpeg_path = path.join(jpeg_folder, row['slice_file_name'])+'.jpg'
            Sound(sound_file_path).export_melgram(jpeg_path)

    def _ensure_image_folders(self):
        print("Creating image folders")
        self._ensure_folder(self.images_path)
        for i in range(1, 11):
            self._ensure_folder(path.join(self.images_path, "fold"+str(i)))

    def _ensure_folder(self, folder):
        try:
            mkdir(folder)
        except FileExistsError:
            pass
