package com.example.csc_knu_mobile_development_project_1

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.material.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import com.example.csc_knu_mobile_development_project_1.ui.theme.Csc_knu_mobile_development_project_1Theme


class MainActivity : ComponentActivity() {
	override fun onCreate(savedInstanceState: Bundle?) {
		super.onCreate(savedInstanceState)
		setContent {
			Csc_knu_mobile_development_project_1Theme {
				Main()
			}
		}
	}
}

@Composable
fun Main() {
	var counter by remember {
		mutableStateOf((1..6).random())
	}
	Column(
		modifier = Modifier
			.fillMaxWidth()
			.fillMaxHeight()
			.clickable { counter--; }
	) {
		if (counter > 0) {
			Main1()
		} else {
			Main2()
		}
	}
}


@Composable
fun Main1() {
	Column(
		modifier = Modifier
			.fillMaxWidth()
			.fillMaxHeight()
	) {
		Text(text = "1")
	}
}

@Composable
fun Main2() {
	Column(
		modifier = Modifier
			.fillMaxWidth()
			.fillMaxHeight()
	) {
		Text(text = "2")
	}
}


@Preview(showBackground = true)
@Composable
fun MainPreview() {
	Main()
}