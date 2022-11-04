package com.example.project_2.core

import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.Column
import androidx.compose.material.Button
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.platform.LocalContext

@Composable
fun ContactsEndsWith(){
    val context = LocalContext.current
    context
    Column() {
        Button(onClick = { /*TODO*/ }) {
            Text("Hey")
        }
    }
}