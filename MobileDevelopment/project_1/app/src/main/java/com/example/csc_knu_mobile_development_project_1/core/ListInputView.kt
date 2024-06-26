package com.example.csc_knu_mobile_development_project_1.core

import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.*
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Delete
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import kotlinx.serialization.json.Json
import kotlinx.serialization.serializer
import java.io.OutputStreamWriter


@Composable
fun SortedListView(list: List<Double>) {
	if (list.isEmpty()) {
		EmptyList()
	} else {
		ListWithSave(list)
	}
}

@Composable
fun ListInputView(
	inputList: List<Double>,
	saveTempList: (list: List<Double>) -> Unit,
	onClickSort: () -> Unit
) {
	var inputList by remember { mutableStateOf(inputList) }
	var numStr by remember { mutableStateOf("") }
	var num = numStr.toDoubleOrNull() ?: 0.0

	WithBottomButton(text = "Sort list!", callback = { onClickSort() }) {
		Column(
			modifier = Modifier
				.fillMaxWidth()
				.padding(horizontal = 10.dp)
				.padding(PaddingValues(bottom = 60.dp, top = 10.dp))
		) {
			TextField(
				label = { Text("New value") },
				modifier = Modifier.fillMaxWidth(),
				keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number),
				singleLine = true,
				value = numStr,
				onValueChange = { numStr = it },
				keyboardActions = KeyboardActions(onDone = {
					if (numStr != "") {
						inputList = inputList + listOf(num)
						saveTempList(inputList)
						numStr = "";
					}
				})
			)
			LazyColumn {
				itemsIndexed(inputList) { i, item ->
					PreviewNumber(
						number = item,
						onDelete = {
							val newList = inputList.toMutableList();
							newList.removeAt(i);
							inputList = newList
							saveTempList(inputList)
						})
				}
			}
		}
	}
}

@Composable
fun PreviewNumber(number: Double, onDelete: (() -> Unit)? = null) {
	Card(
		modifier = Modifier
			.fillMaxWidth()
			.padding(4.dp)
	) {
		Row(horizontalArrangement = Arrangement.SpaceBetween) {
			Text(
				style = MaterialTheme.typography.h6,
				text = number.toString(),
				modifier = Modifier.padding(8.dp)
			)
			onDelete?.let {
				IconButton(
					onClick = it,
				) {
					Icon(Icons.Outlined.Delete, "delete", tint = Color.Gray)
				}
			}
		}
	}
}

@Composable
fun EmptyList() {
	Column(
		modifier = Modifier.fillMaxSize(),
		horizontalAlignment = Alignment.CenterHorizontally,
		verticalArrangement = Arrangement.Center
	) {
		Text(
			"Empty list",
		)
	}
}

@Composable
fun ListWithSave(list: List<Double>) {
	val context = LocalContext.current
	val contentResolver = context.contentResolver

	var uploadLauncher = rememberLauncherForActivityResult(
		contract = ActivityResultContracts.CreateDocument("application/json"),
		onResult = { it ->
			it?.let {
				val writer = OutputStreamWriter(contentResolver.openOutputStream(it))
				writer.write(
					Json.encodeToString(
						Json.serializersModule.serializer(),
						list
					)
				)
				writer.close()
			}
		})
	WithBottomButton(text = "Save list", callback = {
		uploadLauncher.launch("sorted.json")
	}) {
		Column(
			modifier = Modifier
				.fillMaxSize()
				.padding(horizontal = 10.dp)
				.padding(PaddingValues(bottom = 60.dp)),
			horizontalAlignment = Alignment.CenterHorizontally
		) {
			LazyColumn {
				items(list) { item ->
					PreviewNumber(number = item)
				}
			}
		}
	}
}