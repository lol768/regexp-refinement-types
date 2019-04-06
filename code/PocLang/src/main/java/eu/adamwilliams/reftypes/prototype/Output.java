
package eu.adamwilliams.reftypes.prototype;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class Output {

    @SerializedName("MimeType")
    @Expose
    private String mimeType;
    @SerializedName("Value")
    @Expose
    private String value;

    public String getMimeType() {
        return mimeType;
    }

    public void setMimeType(String mimeType) {
        this.mimeType = mimeType;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }

    public Output(String mimeType, String value) {
        this.mimeType = mimeType;
        this.value = value;
    }
}
