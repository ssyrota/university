package com.example.project_2.data

import com.google.android.gms.maps.model.LatLng
import decodePolyline


interface RouteRepository {
    suspend fun getRoute(from: LatLng, to: LatLng): List<LatLng>
}

class NetworkRouteRepository(
    private val mapsApiService: ApiServices
) : RouteRepository {
    override suspend fun getRoute(from: LatLng, to: LatLng): List<LatLng> {
        val polyline = mapsApiService.getDirection(from, to).execute()
            .body()?.routes?.get(0)?.overviewPolyline?.points ?: ""
        return decodePolyline(polyline).map { LatLng(it.latitude, it.longitude) }
    }
}
