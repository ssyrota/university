package com.example.project_2.ui

import androidx.compose.foundation.layout.*
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.ArrowBack
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.currentBackStackEntryAsState
import androidx.navigation.compose.rememberNavController

enum class Screen(val title: String) {
    Main(title = "Compose examples"),
    Contacts(title = "Contacts ends with"),
    Sqlite(title = "Sqlite usage"),
    Maps(title = "Maps with routes")
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ComposeExamplesApp() {
    val navController = rememberNavController()
    val backStackEntry by navController.currentBackStackEntryAsState()
    val currentScreen = Screen.valueOf(
        backStackEntry?.destination?.route ?: Screen.Main.name
    )
    Scaffold(topBar = {
        TopBar(
            title = currentScreen.title,
            canBack = currentScreen != Screen.Main,
            backClick = { navController.navigateUp() }
        )
    }
    ) { innerPadding ->
        NavHost(
            navController = navController,
            startDestination = Screen.Main.name,
            modifier = Modifier.padding(innerPadding)
        ) {
            composable(Screen.Main.name) {
                MainPage(
                    MainPageProps(
                        mapsClick = { navController.navigate(Screen.Maps.name) },
                        contactsClick = {/*TODO*/ },
                        sqliteClick = {/*TODO*/ })
                )
            }
            composable(Screen.Maps.name) {
                MapPage()
            }
        }

    }

}


@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TopBar(
    title: String,
    canBack: Boolean,
    backClick: () -> Unit,
) {
    TopAppBar(
        title = {
            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.Center) {
                Text(
                    text = title,
                    maxLines = 1,
                    overflow = TextOverflow.Ellipsis,
                    style = MaterialTheme.typography.headlineMedium
                )
                if (canBack) {
                    Spacer(modifier = Modifier.padding(10.dp))
                }
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

