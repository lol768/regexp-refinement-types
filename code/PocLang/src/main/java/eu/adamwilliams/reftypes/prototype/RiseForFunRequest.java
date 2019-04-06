package eu.adamwilliams.reftypes.prototype;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;
public class RiseForFunRequest {

    @SerializedName("Version")
    @Expose
    private String version;
    @SerializedName("Source")
    @Expose
    private String source;

    public String getVersion() {
        return version;
    }

    public void setVersion(String version) {
        this.version = version;
    }

    public String getSource() {
        return source;
    }

    public void setSource(String source) {
        this.source = source;
    }

}