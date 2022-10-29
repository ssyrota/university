package com.example.csc_knu_mobile_development_project_1.core

import androidx.compose.foundation.layout.*
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import com.example.csc_knu_mobile_development_project_1.core.data.SortOpsCount
import com.jaikeerthick.composable_graphs.composables.BarGraph


data class SortResultsPageProps(
	val sortedList: List<Double>,
	val viewSortedClick: () -> Unit,
	val sortStats: SortOpsCount
)

@Composable
fun SortResultsPage(props: SortResultsPageProps) {
	WithBottomButton(text = "View sorted list", callback = props.viewSortedClick) {
		Column(
			modifier = Modifier.fillMaxSize(),
			verticalArrangement = Arrangement.Center,
			horizontalAlignment = Alignment.CenterHorizontally
		) {
			if (props.sortedList.isNotEmpty()) {
				OperationsCount(props.sortStats)
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
			} else {
				Text(
					"Empty list is always sorted :D",
				)
			}
		}
	}
}

@Composable
fun OperationsCount(props: SortOpsCount) {
	props.entries.sortedBy { it.key }.map {
		Text(
			"${it.key.printableName} sort: ${it.value} ops",
			modifier = Modifier
				.fillMaxWidth()
				.padding(PaddingValues(bottom = 10.dp)),
			textAlign = TextAlign.Center
		)
	}
}
