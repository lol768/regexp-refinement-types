
package eu.adamwilliams.reftypes.prototype;

import java.util.List;
import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class RiseForFunMetadata {

    @SerializedName("Name")
    @Expose
    private String name;
    @SerializedName("DisplayName")
    @Expose
    private String displayName;
    @SerializedName("Version")
    @Expose
    private double version;
    @SerializedName("TermsOfUseUrl")
    @Expose
    private String termsOfUseUrl;
    @SerializedName("PrivacyUrl")
    @Expose
    private String privacyUrl;
    @SerializedName("Institution")
    @Expose
    private String institution;
    @SerializedName("InstitutionUrl")
    @Expose
    private String institutionUrl;
    @SerializedName("InstitutionImageUrl")
    @Expose
    private String institutionImageUrl;
    @SerializedName("MimeType")
    @Expose
    private String mimeType;
    @SerializedName("SupportsLanguageSyntax")
    @Expose
    private Boolean supportsLanguageSyntax;
    @SerializedName("Title")
    @Expose
    private String title;
    @SerializedName("Description")
    @Expose
    private String description;
    @SerializedName("Question")
    @Expose
    private String question;
    @SerializedName("Url")
    @Expose
    private String url;
    @SerializedName("VideoUrl")
    @Expose
    private String videoUrl;
    @SerializedName("DisableErrorTable")
    @Expose
    private Boolean disableErrorTable;
    @SerializedName("Samples")
    @Expose
    private List<Sample> samples = null;
    @SerializedName("Tutorials")
    @Expose
    private List<Tutorial> tutorials = null;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDisplayName() {
        return displayName;
    }

    public void setDisplayName(String displayName) {
        this.displayName = displayName;
    }

    public double getVersion() {
        return version;
    }

    public void setVersion(double version) {
        this.version = version;
    }

    public String getTermsOfUseUrl() {
        return termsOfUseUrl;
    }

    public void setTermsOfUseUrl(String termsOfUseUrl) {
        this.termsOfUseUrl = termsOfUseUrl;
    }

    public String getPrivacyUrl() {
        return privacyUrl;
    }

    public void setPrivacyUrl(String privacyUrl) {
        this.privacyUrl = privacyUrl;
    }

    public String getInstitution() {
        return institution;
    }

    public void setInstitution(String institution) {
        this.institution = institution;
    }

    public String getInstitutionUrl() {
        return institutionUrl;
    }

    public void setInstitutionUrl(String institutionUrl) {
        this.institutionUrl = institutionUrl;
    }

    public String getInstitutionImageUrl() {
        return institutionImageUrl;
    }

    public void setInstitutionImageUrl(String institutionImageUrl) {
        this.institutionImageUrl = institutionImageUrl;
    }

    public String getMimeType() {
        return mimeType;
    }

    public void setMimeType(String mimeType) {
        this.mimeType = mimeType;
    }

    public Boolean getSupportsLanguageSyntax() {
        return supportsLanguageSyntax;
    }

    public void setSupportsLanguageSyntax(Boolean supportsLanguageSyntax) {
        this.supportsLanguageSyntax = supportsLanguageSyntax;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getQuestion() {
        return question;
    }

    public void setQuestion(String question) {
        this.question = question;
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public String getVideoUrl() {
        return videoUrl;
    }

    public void setVideoUrl(String videoUrl) {
        this.videoUrl = videoUrl;
    }

    public Boolean getDisableErrorTable() {
        return disableErrorTable;
    }

    public void setDisableErrorTable(Boolean disableErrorTable) {
        this.disableErrorTable = disableErrorTable;
    }

    public List<Sample> getSamples() {
        return samples;
    }

    public void setSamples(List<Sample> samples) {
        this.samples = samples;
    }

    public List<Tutorial> getTutorials() {
        return tutorials;
    }

    public void setTutorials(List<Tutorial> tutorials) {
        this.tutorials = tutorials;
    }

}
