package com.example.csc_knu_mobile_development_project_1.core

import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.*
import androidx.compose.material.Button
import androidx.compose.material.ButtonDefaults
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import kotlinx.serialization.json.Json
import kotlinx.serialization.json.double
import kotlinx.serialization.json.jsonArray
import kotlinx.serialization.json.jsonPrimitive
import java.io.BufferedReader
import java.io.InputStreamReader

@Composable
fun MainPage(inputByHandClick: () -> Unit, loadFromFileClick: (l: List<Double>) -> Unit) {
	Row(
		modifier = Modifier
			.fillMaxHeight()
			.fillMaxWidth()
	) {
		Column(
			modifier = Modifier
				.fillMaxWidth()
				.fillMaxHeight(),
			verticalArrangement = Arrangement.Center,
			horizontalAlignment = Alignment.CenterHorizontally
		) {
			LoadFile(loadFromFileClick)
			MainButton("Input list by hand", inputByHandClick)

		}
	}
}

@Composable
fun MainButton(text: String, callback: () -> Unit) {
	Button(
		modifier = Modifier
			.width(300.dp)
			.padding(10.dp),
		onClick = callback, colors = ButtonDefaults.buttonColors(
			backgroundColor = Color.Black,
			contentColor = Color.White
		)
	) {
		Text(
			text = text, style = MaterialTheme.typography.button,
			fontSize = 25.sp
		)
	}
}

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