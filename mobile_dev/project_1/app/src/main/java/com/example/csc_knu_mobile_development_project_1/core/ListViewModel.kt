package com.example.csc_knu_mobile_development_project_1.core

import androidx.lifecycle.ViewModel
import com.example.csc_knu_mobile_development_project_1.core.data.ListUiState
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

class ListViewModel : ViewModel() {
	private val _uiState = MutableStateFlow(ListUiState())
	val uiState: StateFlow<ListUiState> = _uiState.asStateFlow()

}