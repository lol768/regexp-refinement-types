package eu.adamwilliams.reftypes.prototype.api;

import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;

@Retention(RetentionPolicy.RUNTIME)
public @interface RegexRefinement {
    String value();
}
