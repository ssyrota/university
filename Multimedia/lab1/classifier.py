import os
os.environ["KERAS_BACKEND"] = "tensorflow"
import functools
import keras


import keras.layers
import keras.preprocessing
import keras.utils
import tensorflow as tf



class UrbanSoundClassifier:
    def __init__(self, dataset):
        self._dataset = dataset

    def evaluate(self):
        fold_img_folders = self._dataset.folds()
        for i in range(len(fold_img_folders)):
            train = [x for j, x in enumerate(fold_img_folders) if j != i]
            train_datasets = [self._dataset_from_folder(
                fold_img_folders[j]) for j in range(len(train))]
            train_ds = functools.reduce(
                lambda x, y: x.concatenate(y), train_datasets)
            validation_ds = self._dataset_from_folder(fold_img_folders[i])
            model = self._create_model()
            model.fit(train_ds, epochs=10)
            model.evaluate(validation_ds)

    def _dataset_from_folder(self, folder):
        return keras.utils.image_dataset_from_directory(
            folder,
            labels='inferred',
            label_mode='categorical',
            color_mode='rgb',
            batch_size=50,
            image_size=(256, 256),
            seed=2,
        )

    def _create_model(self):
        model = keras.Sequential([
            keras.layers.Conv2D(32, (3, 3), activation='relu',
                                input_shape=(256, 256, 3)),
            keras.layers.MaxPooling2D((2, 2)),
            keras.layers.Conv2D(64, (3, 3), activation='relu'),
            keras.layers.MaxPooling2D((2, 2)),
            keras.layers.Conv2D(128, (3, 3), activation='relu'),
            keras.layers.MaxPooling2D((2, 2)),
            keras.layers.Dropout(0.1),
            keras.layers.Flatten(),
            keras.layers.Dense(128, activation='relu'),
            keras.layers.Dense(10, activation='softmax')
        ])
        model.compile(optimizer='adam',
                      loss='categorical_crossentropy',
                      metrics=['accuracy'])
        return model
