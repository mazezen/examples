package com.mazezen;


import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.interfaces.Claim;
import com.auth0.jwt.interfaces.DecodedJWT;
import org.junit.jupiter.api.Test;

import java.util.Date;
import java.util.HashMap;
import java.util.Map;

public class JwtTest {

    @Test
    public void testGenerateToken() {
        Map<String, Object> map = new HashMap<>();
        map.put("id", 1);
        map.put("username", "zhangsan");
        String token = JWT.create()
                .withClaim("user", map)
                .withExpiresAt(new Date(System.currentTimeMillis()+1000*60*5))
                .sign(Algorithm.HMAC256("mazen"));
        System.out.println(token);
    }

    @Test
    public void testDecodeToken() {
        String token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoxLCJ1c2VybmFtZSI6InpoYW5nc2FuIn0sImV4cCI6MTc0MjE3NzcyMH0.eA15GqaL4Co5BfoP0iotG47iUHvr54YlkB98f114sw4";
        JWTVerifier jwtVerifier = JWT.require(Algorithm.HMAC256("mazen")).build();
        DecodedJWT decodedJWT = jwtVerifier.verify(token);
        Map<String, Claim> claims = decodedJWT.getClaims();
        System.out.println(claims.get("user"));
    }
}
