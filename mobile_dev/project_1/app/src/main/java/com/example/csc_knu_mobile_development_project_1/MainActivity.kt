package com.example.csc_knu_mobile_development_project_1

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import com.example.csc_knu_mobile_development_project_1.core.RabbitSorterApp
import com.example.csc_knu_mobile_development_project_1.ui.theme.Csc_knu_mobile_development_project_1Theme


class MainActivity : ComponentActivity() {
	override fun onCreate(savedInstanceState: Bundle?) {
		super.onCreate(savedInstanceState)
		setContent {
			Csc_knu_mobile_development_project_1Theme {
				RabbitSorterApp()
			}
		}
	}
}
