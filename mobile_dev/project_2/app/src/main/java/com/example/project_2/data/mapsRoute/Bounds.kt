import com.example.project_2.data.mapsRoute.Northeast
import com.example.project_2.data.mapsRoute.Southwest
import com.google.gson.annotations.SerializedName

data class Bounds(
    @SerializedName("northeast")
    var northeast: Northeast?,
    @SerializedName("southwest")
    var southwest: Southwest?
)