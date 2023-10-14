data class Coordinate(val longitude: Double, val latitude: Double)

fun decodePolyline(polyline: String): List<Coordinate> {
    val coordinateChunks: MutableList<MutableList<Int>> = mutableListOf()
    coordinateChunks.add(mutableListOf())

    for (char in polyline.toCharArray()) {
        var value = char.code - 63
        val isLastOfChunk = (value and 0x20) == 0
        value = value and (0x1F)
        coordinateChunks.last().add(value)
        if (isLastOfChunk)
            coordinateChunks.add(mutableListOf())
    }
    coordinateChunks.removeAt(coordinateChunks.lastIndex)
    var coordinates: MutableList<Double> = mutableListOf()
    for (coordinateChunk in coordinateChunks) {
        var coordinate =
            coordinateChunk.mapIndexed { i, chunk -> chunk shl (i * 5) }.reduce { i, j -> i or j }
        if (coordinate and 0x1 > 0)
            coordinate = (coordinate).inv()

        coordinate = coordinate shr 1
        coordinates.add((coordinate).toDouble() / 100000.0)
    }

    val points: MutableList<Coordinate> = mutableListOf()
    var previousX = 0.0
    var previousY = 0.0

    for (i in 0..coordinates.size - 1 step 2) {
        if (coordinates[i] == 0.0 && coordinates[i + 1] == 0.0)
            continue

        previousX += coordinates[i + 1]
        previousY += coordinates[i]

        points.add(Coordinate(round(previousX, 5), round(previousY, 5)))
    }
    return points
}

private fun round(value: Double, precision: Int) =
    (value * Math.pow(10.0, precision.toDouble())).toInt().toDouble() / Math.pow(
        10.0,
        precision.toDouble()
    )
