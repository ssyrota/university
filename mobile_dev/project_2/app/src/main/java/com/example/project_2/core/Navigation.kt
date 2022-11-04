package com.example.project_2.core

import androidx.compose.foundation.layout.*
import androidx.compose.material.*
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.ArrowBack
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.currentBackStackEntryAsState
import androidx.navigation.compose.rememberNavController


enum class Screen(val title: String) {
    Main(title = "Compose examples"),
    Contacts(title = "Contacts filter"),
    Sqlite(title = "Sqlite"),
    Maps(title = "Maps"),
}


@Composable
fun ComposeUsageApp() {
    val navController = rememberNavController()
    val backStackEntry by navController.currentBackStackEntryAsState()
    val currentScreen = Screen.valueOf(
        backStackEntry?.destination?.route ?: Screen.Main.name
    )

    Scaffold(
        topBar = {
            TopBar(
                title = currentScreen.title,
                canBack = currentScreen != Screen.Main,
                backClick = { navController.navigateUp() },
            )
        }
    ) { innerPadding ->
        NavHost(
            navController = navController,
            startDestination = Screen.Main.name,
            modifier = Modifier.padding(innerPadding)
        ) {
            composable(Screen.Main.name) {
            }
            composable(Screen.Contacts.name) {
            }
            composable(Screen.Sqlite.name) {
            }
            composable(Screen.Maps.name) {
            }
        }
    }
}



@Composable
fun TopBar(
    title: String,
    canBack: Boolean,
    backClick: () -> Unit,
) {
    TopAppBar(backgroundColor = Color.Black, contentColor = Color.White,
        title = {
            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.Center) {
                Text(
                    text = title,
                    fontSize = 20.sp,
                    maxLines = 1,
                    overflow = TextOverflow.Ellipsis,
                )
                Spacer(modifier = Modifier.padding(start = 50.dp))
            }
        }, navigationIcon = {
            if (canBack) {
                IconButton(
                    onClick = backClick,
                ) {
                    Icon(Icons.Default.ArrowBack, "back")
                }
            }
        })
}

