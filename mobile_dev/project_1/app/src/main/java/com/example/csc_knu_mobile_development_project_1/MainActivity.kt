package com.example.csc_knu_mobile_development_project_1

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Surface
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.sp
import com.example.csc_knu_mobile_development_project_1.ui.theme.Csc_knu_mobile_development_project_1Theme

class MainActivity : ComponentActivity() {
	override fun onCreate(savedInstanceState: Bundle?) {
		super.onCreate(savedInstanceState)
		setContent {
			Csc_knu_mobile_development_project_1Theme {
				Surface(
					modifier = Modifier.fillMaxSize(),
					color = MaterialTheme.colors.background
				) {
					Author(message = "Serhii Syrota")
				}
			}
		}
	}
}

@Composable
fun Author(message: String) {
	Text(text = message, fontSize = 36.sp)
}

@Preview()
@Composable()
fun DefaultPreview() {
	Author(message = "Serhii Syrota")
}
