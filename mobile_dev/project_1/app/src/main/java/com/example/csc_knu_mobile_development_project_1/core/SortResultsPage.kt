package com.example.csc_knu_mobile_development_project_1.core

import androidx.compose.foundation.layout.*
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import com.jaikeerthick.composable_graphs.composables.BarGraph


data class SortResultsPageProps(val sortedList: List<Double>)

@Composable
fun SortResultsPage(props: SortResultsPageProps) {
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
			Text(props.sortedList.toString())
			BarGraph(
				header = {
					Text(
						"Sorted values histogram",
						modifier = Modifier
							.fillMaxWidth()
							.padding(PaddingValues(bottom = 10.dp)),
						textAlign = TextAlign.Center
					)
				},
				dataList = props.sortedList,
			)
		}
	}
}