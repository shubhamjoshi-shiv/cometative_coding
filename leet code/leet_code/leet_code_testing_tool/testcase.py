def run_test_case():
            s = Solution()
    
            # Example 1
            result = s.trap(height = [0,1,0,2,1,0,1,3,2,1,2,1])
            output_param = 6
            if result != output_param:
                print('expected result -',output_param)
                print('got result -',result)
            else:
                print('tc passed')
            
            # Example 2
            result = s.trap(height = [4,2,0,3,2,5])
            output_param = 9
            if result != output_param:
                print('expected result -',output_param)
                print('got result -',result)
            else:
                print('tc passed')
            
run_test_case()