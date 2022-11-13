package com.example.project_2.ui

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.text.selection.SelectionContainer
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountCircle
import androidx.compose.material.icons.outlined.Email
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp


@Composable
fun AuthorMain() {
    Column(
        Modifier
            .fillMaxWidth()
            .padding(top = 150.dp),
        horizontalAlignment = Alignment.CenterHorizontally,
        verticalArrangement = Arrangement.Center,
    ) {
        Image(
            painter = painterResource(id = com.example.project_2.R.drawable._022_10_05_18_54_33),
            contentDescription = "Author photo",
            modifier = Modifier
                .size(150.dp)
                .align(Alignment.CenterHorizontally)
                .clip(CircleShape)
        )
        Column {
            Text(
                text = "Serhii Syrota",
                style = MaterialTheme.typography.headlineMedium,
                modifier = Modifier
                    .align(Alignment.CenterHorizontally)
                    .padding(top = 13.dp),
                textAlign = TextAlign.Center
            )
            Text(
                text = "Student of faculty CSC group TTP-42",
                style = MaterialTheme.typography.headlineSmall,
                modifier = Modifier
                    .align(Alignment.CenterHorizontally),
                textAlign = TextAlign.Center
            )
        }
    }
}

@Composable
fun AuthorContacts() {
    Column(
        modifier = Modifier
            .fillMaxWidth()
            .padding(top = 40.dp),
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        Column(
            horizontalAlignment = Alignment.Start,
            verticalArrangement = Arrangement.Center
        ) {
            Row {
                Icon(
                    Icons.Outlined.AccountCircle,
                    contentDescription = "github link",
                    modifier = Modifier
                        .size(20.dp)
                        .align(Alignment.CenterVertically)
                )
                SelectionContainer() {
                    Text(
                        style = MaterialTheme.typography.bodyLarge,
                        text = "Github: ssyrota",
                    )
                }
            }
            Row {
                Icon(
                    Icons.Outlined.Email,
                    contentDescription = "email",
                    modifier = Modifier
                        .size(20.dp)
                        .align(Alignment.CenterVertically)
                )
                SelectionContainer() {
                    Text(
                        style = MaterialTheme.typography.bodyLarge,
                        text = "Gmail: serhii_syrota@knu.ua"
                    )
                }
            }
        }
    }
}

@Composable
fun AuthorPage() {
    Column() {
        AuthorMain()
        AuthorContacts()
    }
}