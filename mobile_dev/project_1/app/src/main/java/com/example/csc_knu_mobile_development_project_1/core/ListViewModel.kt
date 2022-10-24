package com.example.csc_knu_mobile_development_project_1.core

import androidx.lifecycle.ViewModel
import com.example.csc_knu_mobile_development_project_1.core.data.ListUiState
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update

class ListViewModel : ViewModel() {
	private val _uiState = MutableStateFlow(ListUiState())
	val uiState: StateFlow<ListUiState> = _uiState.asStateFlow()

	fun setList(items: List<Double>) {
		_uiState.update { currentState ->
			currentState.copy(list = items.toMutableList())
		}
	}
}