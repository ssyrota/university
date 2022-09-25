package com.example.csc_knu_mobile_development_project_1

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.animation.AnimatedVisibility
import androidx.compose.foundation.Image
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material.Card
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.csc_knu_mobile_development_project_1.ui.theme.Csc_knu_mobile_development_project_1Theme


class MainActivity : ComponentActivity() {
	override fun onCreate(savedInstanceState: Bundle?) {
		super.onCreate(savedInstanceState)
		setContent {
			Csc_knu_mobile_development_project_1Theme {
				Author()
			}
		}
	}
}

@Composable
fun Author() {
	Card {
		var expanded by remember { mutableStateOf(false) }
		Column(
			Modifier
				.clickable { expanded = !expanded }) {
			Image(
				painter = painterResource(id = R.drawable._6842028),
				contentDescription = "photo",
				modifier = Modifier
					.padding()
					.size(400.dp)
					.align(Alignment.CenterHorizontally)
					.clip(CircleShape)
			)
			AnimatedVisibility(expanded) {
				Column {
					Text(
						text = "Developed by",
						style = MaterialTheme.typography.h2
					)
					Text(
						text = stringResource(R.string.author_name),
						style = MaterialTheme.typography.h2
					)
				}
			}
		}
	}
}

@Preview(showBackground = true, showSystemUi = true)
@Composable()
fun AuthorPreview() {
	Author()
}