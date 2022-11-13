package com.example.project_2.data

import androidx.annotation.NonNull
import androidx.room.*

@Entity(tableName = "vehicle")
class Vehicle {

    @PrimaryKey(autoGenerate = true)
    @NonNull
    @ColumnInfo(name = "id")
    var id: Int = 0

    val brand: String
    val bodyType: String
    val color: String
    val engineVolume: Double
    val year: Int

    constructor(
        brand: String,
        bodyType: String,
        color: String,
        engineVolume: Double,
        year: Int
    ) {
        this.id = id
        this.brand = brand
        this.bodyType = bodyType
        this.color = color
        this.engineVolume = engineVolume
        this.year = year
    }
}

@Dao
interface VehicleDao {
    @Insert
    fun insertVehicle(data: Vehicle)

    @Query("SELECT * FROM vehicle WHERE bodyType='universal' AND color='red'")
    fun redUniversal(): List<Vehicle>

    @Query("SELECT AVG(engineVolume) FROM vehicle")
    fun engineVolumeAVG(): Double

}