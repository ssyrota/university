package com.example.project_2.data

import android.app.Application
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel

class MainViewModel(application: Application) : ViewModel() {
    private val repository: VehicleRepo
    val redUniversalResult: MutableLiveData<List<Vehicle>>
    val engineVolumeAvg: MutableLiveData<Double>

    init {
        val vehicleDb = VehicleRoomDatabase.instance(application)
        val vehicleDao = vehicleDb.vehicleDao()
        repository = VehicleRepo(vehicleDao)
        redUniversalResult = repository.redUniversal
        engineVolumeAvg = repository.engineAvg
    }


    fun insertCar(data: Vehicle) {
        repository.insert(data)
    }

    fun updateRedUniversal() {
        repository.updateRedUniversal()
    }

    fun updateEngineAvg() {
        repository.updateEngineVolumeAVG()
    }
}
