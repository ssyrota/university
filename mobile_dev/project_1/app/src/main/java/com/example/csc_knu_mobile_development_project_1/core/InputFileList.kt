package com.example.csc_knu_mobile_development_project_1.core

import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.material.Button
import androidx.compose.material.Text
import androidx.compose.runtime.*
import androidx.compose.ui.platform.LocalContext
import kotlinx.serialization.json.Json
import kotlinx.serialization.json.jsonArray
import java.io.BufferedReader
import java.io.InputStreamReader

@Composable
fun FilePicker() {
	var fileContent by remember { mutableStateOf("") }
	val context = LocalContext.current
	val contentResolver = context.contentResolver

	val launcher =
		rememberLauncherForActivityResult(contract = ActivityResultContracts.GetContent()) { uri ->
			uri?.let {
				fileContent =
					BufferedReader(InputStreamReader(contentResolver.openInputStream(it))).readText()
			}
		}
	Button(onClick = {
		launcher.launch("application/json")
	}) {
		Json.parseToJsonElement(fileContent).jsonArray
		Text(text = fileContent)
	}
}