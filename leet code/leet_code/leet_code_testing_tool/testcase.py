def run_test_case():
            s = Solution()
    
            # Example 1
            result = s.encode_and_decode(["neet","code","love","you"])
            output_param = ["neet","code","love","you"]
            if result != output_param:
                print('expected result -',output_param)
                print('got result -',result)
            else:
                print('tc passed')
            
            # Example 2
            result = s.encode_and_decode(["we","say",":","yes"])
            output_param = ["we","say",":","yes"]
            if result != output_param:
                print('expected result -',output_param)
                print('got result -',result)
            else:
                print('tc passed')
            
run_test_case()