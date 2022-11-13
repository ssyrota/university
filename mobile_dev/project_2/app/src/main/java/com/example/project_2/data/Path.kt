package com.example.project_2.data

import com.example.project_2.data.mapsRoute.DirectionResponses
import com.google.android.gms.maps.model.LatLng
import retrofit2.Call
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import retrofit2.http.GET
import retrofit2.http.Query

interface ApiServices {
    @GET("maps/api/directions/json")
    fun getDirection(
        @Query("origin") origin: LatLng,
        @Query("destination") destination: LatLng,
        @Query("key") apiKey: String = "AIzaSyAMjiF3EI0gu7CZ88yHbi7kgsD26r8l9wI"
    ): Call<DirectionResponses>
}

fun apiServices(): ApiServices {
    val retrofit = Retrofit.Builder()
        .addConverterFactory(GsonConverterFactory.create())
        .baseUrl("https://maps.googleapis.com")
        .build()

    return retrofit.create(ApiServices::class.java)
}

object GoogleMapsApi {
    val instance: ApiServices by lazy {
        apiServices()
    }
}