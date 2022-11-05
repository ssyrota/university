package com.example.project_2.core

import androidx.compose.runtime.Composable

data class MainPageProps(
    val mapsClick: () -> Unit,
    val contactsClick: () -> Unit,
    val sqliteClick: () -> Unit,
)

@Composable
fun MainPage(props: MainPageProps) {

}
