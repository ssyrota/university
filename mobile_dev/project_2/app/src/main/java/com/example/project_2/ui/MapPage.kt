package com.example.project_2.ui

import android.content.Context
import android.location.Geocoder
import android.location.Location
import android.os.Looper
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.material3.Button
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.*
import androidx.compose.ui.ExperimentalComposeUiApi
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalFocusManager
import androidx.compose.ui.unit.dp
import com.google.accompanist.permissions.ExperimentalPermissionsApi
import com.google.accompanist.permissions.rememberMultiplePermissionsState
import com.google.android.gms.location.LocationRequest
import com.google.android.gms.location.LocationServices
import com.google.android.gms.maps.model.CameraPosition
import com.google.android.gms.maps.model.LatLng
import com.google.maps.android.compose.*

val SECOND_IN_MILI: Long = 1000

@Composable
fun MapPage() {
    val locationName = remember {
        mutableStateOf("")
    }
    var currentLocation by remember { mutableStateOf(Location(null)) }
    val context = LocalContext.current
    MonitorLocation(context) {
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
                Marker(
                    state = MarkerState(position = pickedPosition),
                    title = locationName.value
                )
            }
            FindLocation(setLocation = {
                locationName.value = it;
                pickedPosition =
                    CoordinatesByName(locationName.value, context)
                cameraPositionState.position = CameraPosition.fromLatLngZoom(pickedPosition, 11f)
            })
        }
    }
}

fun CoordinatesByName(locationName: String, context: Context): LatLng {
    try {
        Geocoder(context).getFromLocationName(locationName, 1).let {
            return LatLng(it?.get(0)?.latitude ?: 0.0, it?.get(0)?.longitude ?: 0.0)
        }
    } catch (e: java.lang.Exception) {
        e.printStackTrace()
    }
    return LatLng(0.0, 0.0)
}

@OptIn(ExperimentalMaterial3Api::class, ExperimentalComposeUiApi::class)
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
        keyboardActions = KeyboardActions { setLocation(place.value); focusManager.clearFocus() }
    )
}

@OptIn(ExperimentalPermissionsApi::class)
@Composable
private fun WithLocation(Child: @Composable () -> Unit) {
    val locationPermissionState = rememberMultiplePermissionsState(
        listOf(
            android.Manifest.permission.ACCESS_COARSE_LOCATION,
            android.Manifest.permission.ACCESS_FINE_LOCATION
        )
    )
    if (locationPermissionState.allPermissionsGranted) {
        Child()
    } else {
        Column {
            val textToShow =
                "Location permission required for this feature to be available. Please grant the Manifest.permission"
            Text(textToShow)
            Button(onClick = { locationPermissionState.launchMultiplePermissionRequest() }) {
                Text("Request permission")
            }
        }
    }
}

fun MonitorLocation(context: Context, callback: (l: Location) -> Unit) {
    LocationServices.getFusedLocationProviderClient(context)
        .requestLocationUpdates(
            LocationRequest.Builder(SECOND_IN_MILI).build(),
            callback,
            Looper.getMainLooper()
        )
}