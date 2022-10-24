package com.example.csc_knu_mobile_development_project_1.core

import androidx.compose.foundation.layout.*
import androidx.compose.material.*
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.ArrowBack
import androidx.compose.material.icons.filled.Info
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.currentBackStackEntryAsState
import androidx.navigation.compose.rememberNavController

enum class Screen(val title: String) {
	Main(title = "Rabbit search"),
	WriteList(title = "Write list"),
	SortResults(title = "Sort results"),
	Author(title = "Author"),
	SortedList(title = "Sorted list")
}


@Composable
fun RabbitSorterApp() {
	val navController = rememberNavController()
	val viewModel: ListViewModel = viewModel()
	val backStackEntry by navController.currentBackStackEntryAsState()
	// Get the name of the current screen
	val currentScreen = Screen.valueOf(
		backStackEntry?.destination?.route ?: Screen.Main.name
	)

	Scaffold(
		topBar = {
			TopBar(
				title = currentScreen.title,
				canBack = currentScreen != Screen.Main,
				backClick = { navController.navigateUp() },
				canAuthor = currentScreen != Screen.Author,
				authorClick = { navController.navigate(Screen.Author.name) },
			)
		}
	) { innerPadding ->
		val uiState by viewModel.uiState.collectAsState()
		NavHost(
			navController = navController,
			startDestination = Screen.Main.name,
			modifier = Modifier.padding(innerPadding)
		) {
			composable(Screen.Main.name) {
				MainPage(loadFromFileClick = {
					navController.navigate(Screen.WriteList.name)
				}, inputByHandClick = {
					navController.navigate(Screen.WriteList.name)
				})
			}
			composable(Screen.Author.name) {
				AuthorPage()
			}
			composable(Screen.SortedList.name) {
				SortedListView(list = uiState.list)
			}
			composable(Screen.WriteList.name) {
				ListInputView(
					list = uiState.list,
					onClickSort = { it -> viewModel.setList(it); navController.navigate(Screen.SortResults.name) })
			}
		}
	}
}

@Composable
@Preview(showBackground = true, showSystemUi = true)
fun RabbitSorterAppPreview() {
	RabbitSorterApp()
}


@Composable
fun TopBar(
	title: String,
	canBack: Boolean,
	backClick: () -> Unit,
	canAuthor: Boolean = true,
	authorClick: () -> Unit
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
				Spacer(modifier = Modifier.padding(10.dp))
			}
		}, navigationIcon = {
			if (canBack) {
				IconButton(
					onClick = backClick,
				) {
					Icon(Icons.Default.ArrowBack, "back")
				}
			}
		}, actions = {
			if (canAuthor) {
				IconButton(
					onClick = authorClick,
				) {
					Icon(Icons.Default.Info, "info")
				}
			} else {
				Spacer(modifier = Modifier.padding(25.dp))
			}
		})
}

