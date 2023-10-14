package com.example.project_2.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import com.example.project_2.data.MainViewModel
import kotlinx.coroutines.launch

@Composable
fun SqliteExample(viewModel: MainViewModel) {
    val coroutineScope = rememberCoroutineScope()
    var engineVolumeAvg by remember { mutableStateOf(viewModel.engineVolumeAvg.value) }
    var redUniversal by remember { mutableStateOf(viewModel.redUniversalResult.value) }
    coroutineScope.launch {
        viewModel.updateEngineAvg()
        viewModel.updateRedUniversal()
        engineVolumeAvg = viewModel.engineVolumeAvg.value ?: 0.0
        redUniversal = viewModel.redUniversalResult.value
    }

    Column {
        Text(
            text = "Engine loud average: ${"%.2f".format(engineVolumeAvg)}",
            textAlign = TextAlign.Center,
            modifier = Modifier.fillMaxWidth()
        )
        LazyColumn(
            Modifier
                .fillMaxSize()
                .padding(16.dp)
        ) {
            item {
                Row(Modifier.background(Color.Gray)) {
                    TableCell(text = "Id", weight = 0.1f)
                    TableCell(text = "Body type", weight = 0.1f)
                    TableCell(text = "Brand", weight = 0.1f)
                    TableCell(text = "Color", weight = 0.1f)
                    TableCell(text = "Engine volume", weight = 0.1f)
                    TableCell(text = "Year", weight = 0.1f)
                }
            }
            items(redUniversal ?: listOf()) {
                Row(Modifier.fillMaxWidth()) {
                    TableCell(text = it.id.toString(), weight = 0.1f)
                    TableCell(text = it.bodyType, weight = 0.1f)
                    TableCell(text = it.brand, weight = 0.1f)
                    TableCell(text = it.color, weight = 0.1f)
                    TableCell(text = it.engineVolume.toString(), weight = 0.1f)
                    TableCell(text = it.year.toString(), weight = 0.1f)
                }
            }
        }

    }
}

@Composable
fun RowScope.TableCell(
    text: String,
    weight: Float
) {
    Text(
        text = text,
        modifier = Modifier
            .border(1.dp, Color.Black)
            .weight(weight)
            .padding(8.dp)
    )
}