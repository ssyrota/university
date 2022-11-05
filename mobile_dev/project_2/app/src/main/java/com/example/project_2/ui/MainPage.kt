package com.example.project_2.ui

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp

data class MainPageProps(
    val mapsClick: () -> Unit,
    val contactsClick: () -> Unit,
    val sqliteClick: () -> Unit,
)

@Composable
fun MainPage(props: MainPageProps) {
    Column(
        modifier = Modifier.fillMaxSize(),
        verticalArrangement = Arrangement.Center,
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        Image(
            painterResource(id = com.example.project_2.R.drawable.jetpack),
            "Jetpack compose logo",
            modifier = Modifier
                .size(200.dp)
                .align(Alignment.CenterHorizontally)
                .clip(CircleShape)
        )
        Spacer(modifier = Modifier.padding(vertical = 50.dp))
        MainButtons(
            MainButtonsProps(
                buttons = listOf(
                    ButtonProps(
                        name = "Contacts",
                        onClick = props.contactsClick
                    ),
                    ButtonProps(
                        name = "Sqlite",
                        onClick = props.sqliteClick
                    ),
                    ButtonProps(
                        name = "Maps",
                        onClick = props.mapsClick
                    )
                )
            )
        )
    }
}

data class ButtonProps(val name: String, val onClick: () -> Unit)
data class MainButtonsProps(val buttons: List<ButtonProps>)

@Composable
fun MainButtons(props: MainButtonsProps) {
    LazyColumn() {
        items(props.buttons) { item ->
            Button(
                onClick = item.onClick,
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(horizontal = 50.dp, vertical = 10.dp)
                    .height(60.dp),
            ) {
                Text(item.name, style = MaterialTheme.typography.bodyLarge)
            }
        }
    }
}