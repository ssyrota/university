package com.example.project_2.data

import androidx.lifecycle.MutableLiveData
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.async
import kotlinx.coroutines.launch

class VehicleRepo(private val dao: VehicleDao) {
    var searchResults = MutableLiveData<List<Vehicle>>()
    var engineAvg = MutableLiveData<Double>()
    private val coroutineScope = CoroutineScope(Dispatchers.Main)

    fun insert(data: Vehicle) {
        coroutineScope.launch(Dispatchers.IO) {
            dao.insertVehicle(data)
        }
    }

    fun updateRedUniversal() =
        coroutineScope.async(Dispatchers.IO) {
            searchResults.value = dao.redUniversal()
        }

    fun updateEngineVolumeAVG() =
        coroutineScope.async(Dispatchers.IO) {
            engineAvg.value = dao.engineVolumeAVG()
        }

}