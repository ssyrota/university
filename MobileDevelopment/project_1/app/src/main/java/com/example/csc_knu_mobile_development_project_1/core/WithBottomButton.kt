package com.example.csc_knu_mobile_development_project_1.core

import androidx.compose.foundation.layout.*
import androidx.compose.material.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

@Composable
fun WithBottomButton(callback: () -> Unit, text: String, child: @Composable () -> Unit) {
	Scaffold(bottomBar = {
		Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.Center) {
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
	}) {
		child()
	}
}