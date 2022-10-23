package com.example.csc_knu_mobile_development_project_1.core

import androidx.compose.foundation.layout.*
import androidx.compose.material.Button
import androidx.compose.material.ButtonDefaults
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

@Composable
fun MainPage(loadFromFileClick: () -> Unit, inputByHandClick: () -> Unit) {
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
			MainButton("Load from file", loadFromFileClick)
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