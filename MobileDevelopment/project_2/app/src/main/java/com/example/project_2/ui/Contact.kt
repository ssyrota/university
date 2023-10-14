package com.example.project_2.ui

import android.Manifest
import android.database.Cursor
import android.provider.ContactsContract
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.runtime.mutableStateListOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalFocusManager
import androidx.compose.ui.unit.dp

@Composable
fun ContactsEndWith() {
    WithContacts {
        val contacts = mutableStateListOf<Contact>()
        val contentRes = LocalContext.current.contentResolver
        val contactsCursor =
            contentRes.query(
                ContactsContract.CommonDataKinds.Phone.CONTENT_URI,
                null,
                null,
                null,
                null
            )
        if (contactsCursor !== null) {
            while (contactsCursor.moveToNext()) {
                contacts.add(contactsCursor.contact())
            }
            ContactsView(contacts.filter { it.number.endsWith("7") })
        }
    }
}


@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun EndsWithForm(setEnds: (v: String) -> Unit) {
    val place = remember {
        mutableStateOf("")
    }
    val focusManager = LocalFocusManager.current
    TextField(
        value = place.value,
        onValueChange = { place.value = it; },
        placeholder = { Text(text = "Enter contacts end with") },
        modifier = Modifier
            .fillMaxWidth()
            .padding(top = 30.dp)
            .padding(horizontal = 20.dp)
            .height(60.dp),
        singleLine = true,
        keyboardActions = KeyboardActions { setEnds(place.value); focusManager.clearFocus() },
    )
}

@Composable
fun ContactsView(contacts: List<Contact>) {
    LazyColumn {
        itemsIndexed(contacts) { i, it ->
            Card(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(4.dp)
            ) {
                Row(horizontalArrangement = Arrangement.SpaceBetween) {
                    Text(
                        style = MaterialTheme.typography.bodyMedium,
                        text = "${it.name} ${it.number}",
                        modifier = Modifier.padding(8.dp)
                    )
                }
            }
        }
    }
}

private fun Cursor.contact(): Contact {
    return Contact(
        this.getString(this.getColumnIndex(ContactsContract.CommonDataKinds.Phone.NUMBER)),
        this.getString(this.getColumnIndex(ContactsContract.CommonDataKinds.Phone.DISPLAY_NAME))
    )
}

data class Contact(val number: String, val name: String)

@Composable
private fun WithContacts(Child: @Composable () -> Unit) {
    WithPermissions(
        listOf(
            Manifest.permission.READ_CONTACTS,
            Manifest.permission.WRITE_CONTACTS
        )
    ) {
        Child()
    }
}