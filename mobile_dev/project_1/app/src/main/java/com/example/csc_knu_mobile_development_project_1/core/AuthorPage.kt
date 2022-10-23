package com.example.csc_knu_mobile_development_project_1.core

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.text.selection.SelectionContainer
import androidx.compose.material.Icon
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Text
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountCircle
import androidx.compose.material.icons.outlined.Email
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
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
			painter = painterResource(id = com.example.csc_knu_mobile_development_project_1.R.drawable._6842028),
			contentDescription = "Author photo",
			modifier = Modifier
				.size(150.dp)
				.align(Alignment.CenterHorizontally)
				.clip(CircleShape)
		)
		Column {
			Text(
				text = stringResource(com.example.csc_knu_mobile_development_project_1.R.string.author_name),
				style = MaterialTheme.typography.h4,
				modifier = Modifier
					.align(Alignment.CenterHorizontally)
					.padding(top = 13.dp),
				textAlign = TextAlign.Center
			)
			Text(
				text = stringResource(com.example.csc_knu_mobile_development_project_1.R.string.author_title),
				style = MaterialTheme.typography.h6,
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
						style = MaterialTheme.typography.h6,
						text = "Github: serhii-syrota",
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
						style = MaterialTheme.typography.h6,
						text = "Gmail: serhii_syrota@knu.ua"
					)
				}
			}
		}
	}
}


// TODO: add back button and header
@Composable
fun AuthorPage() {
	Column() {
		AuthorMain()
		AuthorContacts()
	}
}