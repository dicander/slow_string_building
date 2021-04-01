import java.util.*;

public class SlowBuild {

    public static void main(String[] args) throws Exception{
        int N = Integer.parseInt(args[0]);
        String answer = "";
        for(int i=0; i<N; ++i) {
            answer += i;
        }
        System.out.println(answer);
    }
}
