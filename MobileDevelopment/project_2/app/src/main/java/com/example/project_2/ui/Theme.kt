package com.example.project_2.ui

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.lightColorScheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.graphics.Color

@Composable
fun ExampleTheme(content: @Composable () -> Unit) {
    MaterialTheme(colorScheme = LightColorScheme) {
        content()
    }
}

val LightPrimary = Color(0xFF825500)
val LightOnPrimary = Color(0xFFFFFFFF)
val LightPrimaryContainer = Color(0xFFFFDDAE)
val LightOnPrimaryContainer = Color(0xFF2A1800)
val LightSecondary = Color(0xFF6F5B40)
val LightOnSecondary = Color(0xFFFFFFFF)
val LightSecondaryContainer = Color(0xFFFADEBC)
val LightOnSecondaryContainer = Color(0xFF271904)
val LightTertiary = Color(0xFF516440)
val LightOnTertiary = Color(0xFFFFFFFF)
val LightTertiaryContainer = Color(0xFFD3EABC)
val LightOnTertiaryContainer = Color(0xFF102004)
val LightError = Color(0xFFBA1B1B)
val LightErrorContainer = Color(0xFFFFDAD4)
val LightOnError = Color(0xFFFFFFFF)
val LightOnErrorContainer = Color(0xFF410001)
val LightBackground = Color(0xFFFCFCFC)
val LightOnBackground = Color(0xFF1F1B16)
val LightSurface = Color(0xFFFCFCFC)
val LightOnSurface = Color(0xFF1F1B16)
val LightSurfaceVariant = Color(0xFFF0E0CF)
val LightOnSurfaceVariant = Color(0xFF4F4539)
val LightOutline = Color(0xFF817567)
val LightInverseOnSurface = Color(0xFFF9EFE6)
val LightInverseSurface = Color(0xFF34302A)
val LightPrimaryInverse = Color(0xFFFFB945)

private val LightColorScheme = lightColorScheme(
    primary = LightPrimary,
    onPrimary = LightOnPrimary,
    primaryContainer = LightPrimaryContainer,
    onPrimaryContainer = LightOnPrimaryContainer,
    inversePrimary = LightPrimaryInverse,
    secondary = LightSecondary,
    onSecondary = LightOnSecondary,
    secondaryContainer = LightSecondaryContainer,
    onSecondaryContainer = LightOnSecondaryContainer,
    tertiary = LightTertiary,
    onTertiary = LightOnTertiary,
    tertiaryContainer = LightTertiaryContainer,
    onTertiaryContainer = LightOnTertiaryContainer,
    error = LightError,
    onError = LightOnError,
    errorContainer = LightErrorContainer,
    onErrorContainer = LightOnErrorContainer,
    background = LightBackground,
    onBackground = LightOnBackground,
    surface = LightSurface,
    onSurface = LightOnSurface,
    inverseSurface = LightInverseSurface,
    inverseOnSurface = LightInverseOnSurface,
    surfaceVariant = LightSurfaceVariant,
    onSurfaceVariant = LightOnSurfaceVariant,
    outline = LightOutline
)

