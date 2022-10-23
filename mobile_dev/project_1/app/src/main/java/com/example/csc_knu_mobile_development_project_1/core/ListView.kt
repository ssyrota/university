package com.example.csc_knu_mobile_development_project_1.core

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material.*
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Delete
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import com.example.csc_knu_mobile_development_project_1.core.data.SortedList


@Composable
fun SortedListView(list: List<Double>) {
	LazyColumn {
		items(SortedList(list).insertionSort()) { item ->
			PreviewNumber(number = item)
		}
	}
}

@Composable
fun ListInputView(list: List<Double>) {
	LazyColumn {
		items(SortedList(list).insertionSort()) { item ->
			PreviewNumber(number = item)
		}
	}
}

@Composable
fun PreviewNumber(number: Double, onDelete: (() -> Unit)? = null) {
	Card(
		modifier = Modifier
			.fillMaxWidth()
			.padding(4.dp)
	) {
		Row(horizontalArrangement = Arrangement.SpaceBetween) {
			Text(
				style = MaterialTheme.typography.h6,
				text = number.toString(),
				modifier = Modifier.padding(8.dp)
			)
			onDelete?.let {
				IconButton(
					onClick = it,
				) {
					Icon(Icons.Outlined.Delete, "delete", tint = Color.Gray)
				}
			}
		}
	}
}
