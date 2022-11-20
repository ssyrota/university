package com.example.project_2.data

import android.util.Log
import com.google.android.gms.maps.model.LatLng
import decodePolyline


interface RouteRepository {
    suspend fun getRoute(from: LatLng, to: LatLng): List<LatLng>
}

class NetworkRouteRepository(
    private val mapsApiService: ApiServices
) : RouteRepository {
    override suspend fun getRoute(from: LatLng, to: LatLng): List<LatLng> {
        try {
            val routes =
                mapsApiService.getDirection(from.toMapsString(), to.toMapsString()).execute()
                    .body()?.routes ?: listOf()
            Log.d("ERR", routes.joinToString { e -> e.toString() })
            val polyline =
                if (routes.size > 0) routes.get(0)?.overviewPolyline?.points ?: "" else ""
            return decodePolyline(polyline).map { LatLng(it.latitude, it.longitude) }
        } catch (e: Exception) {
            Log.e("ERR", e.toString())
            return listOf()
        }
    }
}

object NetworkRouteRepo {
    val instance: NetworkRouteRepository by lazy {
        NetworkRouteRepository(GoogleMapsApi.instance)
    }
}

fun LatLng.toMapsString(): String {
    return "${this.latitude},${this.longitude}"
}