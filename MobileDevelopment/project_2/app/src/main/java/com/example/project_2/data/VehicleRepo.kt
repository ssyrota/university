package com.example.project_2.data

import androidx.lifecycle.MutableLiveData
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.async
import kotlinx.coroutines.launch

class VehicleRepo(private val dao: VehicleDao) {
    var redUniversal = MutableLiveData<List<Vehicle>>()
    var engineAvg = MutableLiveData<Double>()
    private val coroutineScope = CoroutineScope(Dispatchers.Main)

    fun insert(data: Vehicle) =
        coroutineScope.launch(Dispatchers.IO) {
            dao.insertVehicle(data)
        }

    fun updateRedUniversal() =
        coroutineScope.launch(Dispatchers.Main) {
            redUniversal.value = redUniversal().await()
        }

    fun updateEngineVolumeAVG() =
        coroutineScope.launch(Dispatchers.Main) {
            engineAvg.value = engineVolumeAvg().await()
        }

    fun engineVolumeAvg() = coroutineScope.async(Dispatchers.IO) {
        return@async dao.engineVolumeAVG()
    }

    fun redUniversal() = coroutineScope.async(Dispatchers.IO) {
        return@async dao.redUniversal()
    }
}