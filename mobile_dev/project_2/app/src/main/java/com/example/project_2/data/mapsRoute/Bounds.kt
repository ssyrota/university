import com.google.gson.annotations.SerializedName
import com.utsman.samplegooglemapsdirection.kotlin.model.Northeast
import com.utsman.samplegooglemapsdirection.kotlin.model.Southwest

data class Bounds(
    @SerializedName("northeast")
    var northeast: Northeast?,
    @SerializedName("southwest")
    var southwest: Southwest?
)