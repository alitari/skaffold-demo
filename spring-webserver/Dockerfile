FROM maven:3.6.3-adoptopenjdk-15 AS MAVEN_BUILD
COPY pom.xml /build/
WORKDIR /build/
RUN mvn verify clean --fail-never
COPY src /build/src/
RUN mvn package -Dmaven.test.skip=true
FROM adoptopenjdk:15-openj9
WORKDIR /app
COPY --from=MAVEN_BUILD /build/target/webserver-0.0.1-SNAPSHOT.jar /app/
ENTRYPOINT ["java", "-jar", "webserver-0.0.1-SNAPSHOT.jar"]