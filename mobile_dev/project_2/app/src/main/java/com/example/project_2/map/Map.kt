package com.example.project_2.map

import android.content.ActivityNotFoundException
import android.content.Context
import android.content.Intent
import android.net.Uri
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.material.Button
import androidx.compose.material.Text
import androidx.compose.material.TextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

@Composable
fun MapPage() {
    val context = LocalContext.current
    val destinationLocation = remember {
        mutableStateOf("")
    }
    Column(
        modifier = Modifier
            .fillMaxHeight()
            .fillMaxWidth()
            .background(Color.White)
    ) {
        Spacer(modifier = Modifier.height(20.dp))
        Text(
            text = "Draw Route on Google Maps in Android",
            textAlign = TextAlign.Center,
            color = Color.Green,
            fontWeight = FontWeight.Bold,
            modifier = Modifier
                .padding(10.dp)
                .fillMaxWidth()
        )
        TextField(
            value = destinationLocation.value,
            onValueChange = { destinationLocation.value = it },
            placeholder = { Text(text = "Enter your destination location") },
            modifier = Modifier.padding(15.dp).fillMaxWidth(),
            textStyle = TextStyle(color = Color.Black, fontSize = 15.sp),
            singleLine = true,
        )
        Spacer(modifier = Modifier.height(20.dp))
        Button(
            onClick = {
                drawTrack( destinationLocation.value, context)
            },
            modifier = Modifier.padding(10.dp).fillMaxWidth()
        ) {
            Text(
                text = "Draw Route on Google Maps",
                color = Color.White,
                textAlign = TextAlign.Center
            )
        }
    }
}

fun drawTrack( destination: String, context: Context) {
    try {
        val uri: Uri = Uri.parse("https://www.google.co.in/maps/dir//$destination")
        val i = Intent(Intent.ACTION_VIEW, uri)
        i.setPackage("com.google.android.apps.maps")
        i.flags = Intent.FLAG_ACTIVITY_NEW_TASK
        context.startActivity(i)
    } catch (e: ActivityNotFoundException) {
        val uri: Uri = Uri.parse("https://play.google.com/store/apps/details?id=com.google.android.apps.maps")
        val i = Intent(Intent.ACTION_VIEW, uri)
        i.flags = Intent.FLAG_ACTIVITY_NEW_TASK
        context.startActivity(i)
    }
}