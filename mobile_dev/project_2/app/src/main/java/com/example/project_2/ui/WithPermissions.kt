package com.example.project_2.ui

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.runtime.SideEffect
import com.google.accompanist.permissions.ExperimentalPermissionsApi
import com.google.accompanist.permissions.rememberMultiplePermissionsState


@OptIn(ExperimentalPermissionsApi::class)
@Composable
fun WithPermissions(
    permissions: List<String>,
    Child: @Composable () -> Unit
) {
    val locationPermissionState = rememberMultiplePermissionsState(
        permissions
    )
    if (locationPermissionState.allPermissionsGranted) {
        Child()
    } else {
        Column {
            SideEffect {
                locationPermissionState.launchMultiplePermissionRequest()
            }
        }
    }
}