
package eu.adamwilliams.reftypes.prototype;

import java.util.List;
import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class Tutorial {

    @SerializedName("Name")
    @Expose
    private String name;
    @SerializedName("Source")
    @Expose
    private String source;
    @SerializedName("Samples")
    @Expose
    private List<Sample> samples = null;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getSource() {
        return source;
    }

    public void setSource(String source) {
        this.source = source;
    }

    public List<Sample> getSamples() {
        return samples;
    }

    public void setSamples(List<Sample> samples) {
        this.samples = samples;
    }

}
