import java.util.*;

public class FastBuild {

    public static void main(String[] args) throws Exception{
        int N = Integer.parseInt(args[0]);
        StringBuilder answer = new StringBuilder("");
        for(int i=0; i<N; ++i) {
            answer.append(""+i);
        }
        System.out.println(answer);
    }
}
