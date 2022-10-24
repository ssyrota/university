package com.example.csc_knu_mobile_development_project_1.core

import androidx.compose.foundation.layout.*
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier


@Composable
fun SortResultsPage() {
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
			Text("Load from file")
			Text("Load from file")
		}
	}
}