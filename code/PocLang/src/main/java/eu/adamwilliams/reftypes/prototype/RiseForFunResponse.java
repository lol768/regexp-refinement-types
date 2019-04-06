
package eu.adamwilliams.reftypes.prototype;

import java.util.List;
import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class RiseForFunResponse {

    @SerializedName("Version")
    @Expose
    private double version;
    @SerializedName("Outputs")
    @Expose
    private List<Output> outputs = null;

    public double getVersion() {
        return version;
    }

    public void setVersion(double version) {
        this.version = version;
    }

    public List<Output> getOutputs() {
        return outputs;
    }

    public void setOutputs(List<Output> outputs) {
        this.outputs = outputs;
    }

}
