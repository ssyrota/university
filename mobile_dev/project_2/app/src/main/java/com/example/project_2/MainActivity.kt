package com.example.project_2

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import com.example.project_2.core.ComposeUsageApp
import com.example.project_2.ui.theme.Project_2Theme


class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            Project_2Theme {
                ComposeUsageApp()
            }
        }
    }
}
