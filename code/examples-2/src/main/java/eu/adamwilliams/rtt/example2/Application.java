package eu.adamwilliams.rtt.example2;

import eu.adamwilliams.reftypes.prototype.ast.BodyEvaluator;
import eu.adamwilliams.reftypes.prototype.ast.FunctionDeclaration;
import eu.adamwilliams.reftypes.prototype.parser.PocLang;
import eu.adamwilliams.reftypes.prototype.parser.PocLex;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Scanner;

public class Application {
    public static void main(String[] args) {
        System.out.println("Please enter your university ID and I will build an LDAP query:");
        Scanner scan = new Scanner(System.in);
        String name = scan.nextLine();
        exec(name);
    }

    private static void exec(String name) {
        PocLex lexer = null;
        try {
            lexer = new PocLex(CharStreams.fromFileName("src/rtt/ldap.txt"));
            CommonTokenStream tokens = new CommonTokenStream(lexer);
            PocLang parser = new PocLang(tokens);
            PocLang.ProgramContext program = parser.program();
            eu.adamwilliams.reftypes.prototype.Application app = new eu.adamwilliams.reftypes.prototype.Application();
            FunctionDeclaration func = app.doTypeChecks(program).getFunctionTable().getFunctionByIdentifier("QueryUser");

            BodyEvaluator.evaluateFunction(func, name);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
