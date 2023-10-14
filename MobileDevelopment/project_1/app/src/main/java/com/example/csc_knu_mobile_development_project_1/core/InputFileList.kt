package com.example.csc_knu_mobile_development_project_1.core

import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.runtime.Composable
import androidx.compose.ui.platform.LocalContext
import kotlinx.serialization.json.Json
import kotlinx.serialization.json.double
import kotlinx.serialization.json.jsonArray
import kotlinx.serialization.json.jsonPrimitive
import java.io.BufferedReader
import java.io.InputStreamReader


@Composable
fun LoadFile(inputByFileClick: (l: List<Double>) -> Unit) {
	val context = LocalContext.current
	val contentResolver = context.contentResolver
	val launcher =
		rememberLauncherForActivityResult(contract = ActivityResultContracts.GetContent()) { uri ->
			uri?.let {
				val content = BufferedReader(
					InputStreamReader(
						contentResolver.openInputStream(
							it
						)
					)
				).readText()
				inputByFileClick(
					jsonToListDouble(
						content
					)
				)
			}
		}
	MainButton("Load from file") {
		launcher.launch("application/json")
	}
}

fun jsonToListDouble(t: String): List<Double> {
	return Json.parseToJsonElement(t).jsonArray.map {
		it.jsonPrimitive.double
	}
}