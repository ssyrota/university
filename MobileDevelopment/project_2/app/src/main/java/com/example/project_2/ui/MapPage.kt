package com.example.project_2.ui

import android.content.Context
import android.location.Geocoder
import android.location.Location
import android.os.Looper
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalFocusManager
import androidx.compose.ui.unit.dp
import com.example.project_2.data.NetworkRouteRepo
import com.google.android.gms.location.LocationRequest
import com.google.android.gms.location.LocationServices
import com.google.android.gms.maps.model.CameraPosition
import com.google.android.gms.maps.model.LatLng
import com.google.maps.android.compose.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

const val SECOND_IN_MILLISECOND: Long = 1000

typealias UpdateRoute = (from: LatLng, to: LatLng) -> Unit

@Composable
fun MapPage() {
    val coroutineScope = rememberCoroutineScope()
    var route by remember { mutableStateOf(mutableListOf<LatLng>()) }
    val updateRoute: UpdateRoute = { from, to ->
        coroutineScope.launch {
            withContext(Dispatchers.IO) {
                val result = NetworkRouteRepo.instance.getRoute(from, to)
                route = result as MutableList<LatLng>
            }
        }
    }
    MapScreen(route, updateRoute)
}

@Composable
fun MapScreen(routes: List<LatLng>, updateRoute: UpdateRoute) {
    val locationName = remember {
        mutableStateOf("")
    }
    var currentLocation by remember { mutableStateOf(Location(null)) }
    val context = LocalContext.current
    monitorLocation(context) {
        currentLocation = it
    }
    var pickedPosition by remember {
        mutableStateOf(LatLng(0.0, 0.0))
    }
    val cameraPositionState = rememberCameraPositionState {
        position = CameraPosition.fromLatLngZoom(pickedPosition, 0f)
    }
    WithLocation {
        Column(
            modifier = Modifier.fillMaxSize(),
            verticalArrangement = Arrangement.Center
        ) {
            GoogleMap(
                cameraPositionState = cameraPositionState,
                properties = MapProperties(isMyLocationEnabled = true),
                modifier = Modifier.height(600.dp)
            ) {
                Polyline(points = routes)
                Marker(
                    state = MarkerState(position = pickedPosition),
                    title = locationName.value
                )
            }
            FindLocation(setLocation = {
                locationName.value = it;
                pickedPosition =
                    coordinatesByName(locationName.value, context)
                updateRoute(currentLocation.latLng(), pickedPosition)
                cameraPositionState.position = CameraPosition.fromLatLngZoom(pickedPosition, 11f)
            })
        }
    }
}

fun Location.latLng(): LatLng {
    return LatLng(this.latitude, this.longitude)
}

fun coordinatesByName(locationName: String, context: Context): LatLng {
    try {
        Geocoder(context).getFromLocationName(locationName, 1).let {
            return LatLng(it?.get(0)?.latitude ?: 0.0, it?.get(0)?.longitude ?: 0.0)
        }
    } catch (e: java.lang.Exception) {
        e.printStackTrace()
    }
    return LatLng(0.0, 0.0)
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun FindLocation(setLocation: (v: String) -> Unit) {
    val place = remember {
        mutableStateOf("")
    }
    val focusManager = LocalFocusManager.current
    TextField(
        value = place.value,
        onValueChange = { place.value = it; },
        placeholder = { Text(text = "Enter your location to search") },
        modifier = Modifier
            .fillMaxWidth()
            .padding(top = 30.dp)
            .padding(horizontal = 20.dp)
            .height(60.dp),
        singleLine = true,
        keyboardActions = KeyboardActions { focusManager.clearFocus(); setLocation(place.value); },
    )
}

@Composable
private fun WithLocation(Child: @Composable () -> Unit) {
    WithPermissions(
        listOf(
            android.Manifest.permission.ACCESS_COARSE_LOCATION,
            android.Manifest.permission.ACCESS_FINE_LOCATION
        )
    ) {
        Child()
    }
}

fun monitorLocation(context: Context, callback: (l: Location) -> Unit) {
    LocationServices.getFusedLocationProviderClient(context)
        .requestLocationUpdates(
            LocationRequest.Builder(SECOND_IN_MILLISECOND).build(),
            callback,
            Looper.getMainLooper()
        )
}